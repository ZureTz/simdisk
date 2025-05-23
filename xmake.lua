-- Generate compile_commands.json for clangd every time the project is built
add_rules("plugin.compile_commands.autoupdate", {lsp = "clangd"})
-- Set c++ code standard: c++20
set_languages("c++20")
add_requires("drogon")

target("simdisk")
  set_kind("binary")
  add_includedirs("include")
  add_files("src/controller/*.cc")
  add_files("src/main.cc")
  add_packages("drogon")

