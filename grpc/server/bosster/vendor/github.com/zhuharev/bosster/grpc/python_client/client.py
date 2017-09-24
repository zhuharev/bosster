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

  target_vk = bosster_pb2.PostJob(type=bosster_pb2.VK,
  social_id = "39816168",
  social_token = "4722691fb64cd18d27bf2af4e4c355c26d902c32a0be0e3eb71932e403187566490406e05bc31437f3e3b",
  )

  post = bosster_pb2.Post(message='hello from bosster',image_urls=['https://pp.userapi.com/c639320/v639320146/51c82/NLu2r3Q3sJw.jpg']);

  req = bosster_pb2.PostRequest(post=post,
  sync=True,
  targets=[target_vk])

  response = stub.Post(req)

  print("client received: " + str(response.jobs[0].status))
  print("{0}".format(response.jobs[0].status == bosster_pb2.ENQUEUED))



if __name__ == '__main__':
  run()
