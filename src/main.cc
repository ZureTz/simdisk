#include <drogon/HttpAppFramework.h>
#include <trantor/utils/Logger.h>

using namespace drogon;

int main() {
  // Version of drogon
  LOG_DEBUG << "Drogon version: " << getVersion();
  // Current working directory is the root of the project
  LOG_DEBUG << "Current working directory: " << std::filesystem::current_path();

  // Load config file
  app().loadConfigFile("config.json");

  // Run HTTP framework,the method will block in the internal event loop
  app().run();
  return 0;
}