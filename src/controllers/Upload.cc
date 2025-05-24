#include "controllers/Upload.h"
#include <drogon/MultiPart.h>
#include <iostream>
#include <trantor/utils/Logger.h>

using namespace api;

// Add definition of your processing function here

void Upload::upload(const HttpRequestPtr &req,
                    std::function<void(const HttpResponsePtr &)> &&callback) {
  LOG_DEBUG << "Upload::upload called";

  std::cout << "Request body: " << req->body() << std::endl;

  MultiPartParser fileUploadParser;
  const bool parseResult = fileUploadParser.parse(req);
  if (!parseResult) {
    LOG_ERROR << "Failed to parse multipart request";

    Json::Value errorResponse;
    errorResponse["error"] = "Failed to parse multipart request";

    auto resp = HttpResponse::newHttpJsonResponse(errorResponse);
    resp->setStatusCode(k400BadRequest);

    callback(resp);
    return;
  }

  const auto params = fileUploadParser.getParameters();
  LOG_DEBUG << "Multipart request parsed successfully, parameters count: "
            << params.size();
  for (const auto &param : params) {
    LOG_DEBUG << "Parameter: " << param.first << " = " << param.second;
  }

  // Check if the request contains files
  if (fileUploadParser.getFilesMap().size() == 0) {
    LOG_ERROR << "No files found in the request";

    Json::Value errorResponse;
    errorResponse["error"] = "No files found in the request";

    auto resp = HttpResponse::newHttpJsonResponse(errorResponse);
    resp->setStatusCode(k400BadRequest);

    callback(resp);
    return;
  }

  // Get the files
  const auto &files = fileUploadParser.getFiles();
  for (const auto &file : files) {
    LOG_DEBUG << "File name: " << file.getFileName();
    // Save the file to disk or process it as needed
    // file.saveAs(savePath + "/" + file.getFileName());
    file.save();
  }

  // Log the file name and size
  LOG_DEBUG << "File body size: " << req->body().size() << " bytes";

  Json::Value ret;
  ret["result"] = "ok";
  auto resp = HttpResponse::newHttpJsonResponse(ret);
  callback(resp);
}
