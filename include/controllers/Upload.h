#pragma once

#include <drogon/HttpController.h>

using namespace drogon;

namespace api {
class Upload : public HttpController<Upload> {
public:
  METHOD_LIST_BEGIN
  // use METHOD_ADD to add your custom processing function here;
  METHOD_ADD(Upload::upload, "/single", Post);
  METHOD_LIST_END
  // your declaration of processing function maybe like this:
  void upload(const HttpRequestPtr &req,
             std::function<void(const HttpResponsePtr &)> &&callback);
};
} // namespace api
