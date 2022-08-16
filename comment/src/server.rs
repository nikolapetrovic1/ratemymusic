use tonic::{transport::Server, Request, Response, Status};
use comment::comment_server::{Comment, CommentServer};
use comment::{CommentRequest,CommentResponse};

pub mod comment {
    tonic::include_proto!("comment");
}
#[derive(Debug,Default)]
pub struct CommentService{

}
#[tonic::async_trait]
impl Comment for CommentService{
        async fn create_comment(&self,request: Request<CommentRequest>,) -> Result<Response<CommentResponse>, Status> {
            println!("Got a request: {:?}", request);

            let req = request.into_inner();
    
            let reply = CommentResponse {
                status: format!("Sent {} BTC to", 200).into(),
            };
            Ok(Response::new(reply))
        }
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50055".parse()?;
    let comment_service = CommentService::default();

    Server::builder()
        .add_service(CommentServer::new(comment_service))
        .serve(addr)
        .await?;

    Ok(())
}