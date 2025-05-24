#include "controllers/Upload.h"

using namespace api;

// Add definition of your processing function here

void Upload::upload(const HttpRequestPtr &req,
                    std::function<void(const HttpResponsePtr &)> &&callback) {
  LOG_DEBUG << "Upload::upload called";

  // If body is empty, return an error
  if (req->body().empty()) {
    LOG_ERROR << "File not provided in the request";

    Json::Value errorResponse;
    errorResponse["error"] = "File not provided";

    auto resp = HttpResponse::newHttpJsonResponse(errorResponse);
    resp->setStatusCode(k400BadRequest);

    callback(resp);
    return;
  }

  // Log the file name and size
  LOG_DEBUG << "File body size: " << req->body().size() << " bytes";

  Json::Value ret;
  ret["result"] = "ok";
  auto resp = HttpResponse::newHttpJsonResponse(ret);
  callback(resp);
}
