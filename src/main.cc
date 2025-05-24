#include <iostream>

#include <drogon/HttpAppFramework.h>
#include <trantor/utils/Logger.h>

using namespace drogon;

int main() {
  // Version of drogon
  std::cout << "Drogon version: " << getVersion() << std::endl;
  // Current working directory is the root of the project
  std::cout << "Current working directory: " << std::filesystem::current_path()
            << std::endl;

  // Load config file
  app().loadConfigFile("config.json");

  // Run HTTP framework,the method will block in the internal event loop
  app().run();
  return 0;
}