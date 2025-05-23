#include <iostream>

#include <drogon/HttpAppFramework.h>

int main() {
  // Current working directory is the root of the project
  std::cout << "Current working directory: " << std::filesystem::current_path()
            << std::endl;

  // Load config file
  drogon::app().loadConfigFile("config.json");
  // Run HTTP framework,the method will block in the internal event loop
  drogon::app().run();
  return 0;
}