package main

/*
#cgo pkg-config: luajit

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

const char* cached_script;

void LuaInit() {
}

const char* LuaCall(int cached) {

  lua_State *L = luaL_newstate();

  luaL_openlibs(L);

  if( cached == 0) {
    luaL_dofile(L, "script.lua");
  } else {
    luaL_dostring(L, cached_script);
  }

  const char* input = "hello from lua\n";

  lua_getglobal(L, "dispatch");
  lua_pushstring(L, input);
  lua_pcall(L, 1, 1, 0);

  const char* lua_output = lua_tostring(L, -1);

  char* output = strdup(lua_output);

  lua_close(L);

  return output;
}
*/
import "C"

import(
  "io/ioutil"
  "net/http"
  "unsafe"
)

// LuaInit does nothing!
func LuaInit() {
  C.LuaInit()
}

// LuaCall calls the script, if cached is set to 0, Lua will read the script from disk.
func LuaCall(cached int) *C.char {
  return C.LuaCall(C.int(cached))
}

// LuaCacheScript will keep the script in memory.
func LuaCacheScript() {
  data, _ := ioutil.ReadFile("script.lua")
  C.cached_script = C.CString(string(data))
}

// NotCachedHandler calls the Lua script with cached = 0
func NotCachedHandler(w http.ResponseWriter, req *http.Request) {
  var cOutput *C.char
  cOutput = LuaCall(0)

  output := C.GoString(cOutput)

  C.free(unsafe.Pointer(cOutput))

  w.Write([]byte(output))
}

// CachedHandler calls the Lua script with cached = 1
func CachedHandler(w http.ResponseWriter, req *http.Request) {
  var cOutput *C.char
  cOutput = LuaCall(1)

  output := C.GoString(cOutput)

  C.free(unsafe.Pointer(cOutput))

  w.Write([]byte(output))
}

// GoHandler doesn't call Lua!
func GoHandler(w http.ResponseWriter, req *http.Request) {
  w.Write([]byte("hello from go\n"))
}

func main() {
  LuaCacheScript()
  LuaInit()

  http.HandleFunc("/not_cached", NotCachedHandler)
  http.HandleFunc("/cached", CachedHandler)
  http.HandleFunc("/go", GoHandler)

  http.ListenAndServe(":5000", nil)
}
