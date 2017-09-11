from __future__ import print_function

import grpc

import bosster_pb2
import bosster_pb2_grpc


def run():
  channel = grpc.insecure_channel('localhost:50051')
  stub = bosster_pb2_grpc.PosterStub(channel)

  post = bosster_pb2.Post(message='hello',image_urls=['https://...']);
  response = stub.Post(bosster_pb2.PostRequest(post=post, async=True,
  social_token="token", social_login="login", social_password="pass"))
  
  print("client received: " + str(response.jobs[0].status))
  print("{0}".format(response.jobs[0].status == bosster_pb2.ENQUEUED))

  response = stub.Post(bosster_pb2.PostRequest(post=post, async=False))
  print("client received: " + str(response.jobs[0].status))
  print("{0}".format(response.jobs[0].status == bosster_pb2.ENQUEUED))


if __name__ == '__main__':
  run()
