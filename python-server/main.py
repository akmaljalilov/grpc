import logging
from concurrent import futures

import grpc

import recommendations_pb2 as recommendations_pb2
import recommendations_pb2_grpc as recommendations_pb2_grpc


class Recommendation(recommendations_pb2_grpc.RecommendationsServicer):
    def Recommend(self, request, context):
        return recommendations_pb2.RecommendationResponse(recommendations=[
            recommendations_pb2.BookRecommendation(id=1, title="The Maltese Falcon"),
            recommendations_pb2.BookRecommendation(id=2, title="Murder on the Orient Express"),
            recommendations_pb2.BookRecommendation(id=3, title="The Hound of the Baskervilles"),
        ])

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    recommendations_pb2_grpc.add_RecommendationsServicer_to_server(Recommendation(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
