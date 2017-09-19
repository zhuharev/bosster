from __future__ import print_function

import grpc

import bosster_pb2
import bosster_pb2_grpc


def run():
  channel = grpc.insecure_channel('localhost:2020')
  stub = bosster_pb2_grpc.PosterStub(channel)

  target_fb = bosster_pb2.PostJob(type=bosster_pb2.FACEBOOK,
  # id группы
  social_id = "24234234",
  # токен админа, от его имени будет опубликогвана запись
  social_token = "",

  social_app_id = "",
  social_app_secret = "",
  )

  post = bosster_pb2.Post(message='hello',image_urls=['https://...']);

  req = bosster_pb2.PostRequest(post=post,
  sync=False,
  targets=[target_fb])

  response = stub.Post(req)

  print("client received: " + str(response.jobs[0].status))
  print("{0}".format(response.jobs[0].status == bosster_pb2.ENQUEUED))



if __name__ == '__main__':
  run()
