#[macro_use]
extern crate diesel;
extern crate dotenv;

use db::establish_connection;
use models::map_comment_to_comment_request;
use tonic::{transport::Server, Request, Response, Status};
use comment::comments_server::{Comments,CommentsServer};
use comment::{CommentRequest,CommentResponse,AllCommentsResponse,IdRequest};
use crate::comment_ops::{create_comment,update_comment,get_by_review};



mod comment_ops;
mod models;
mod db;
mod schema;
pub mod comment {
    tonic::include_proto!("comment");
}
#[derive(Debug,Default)]
pub struct CommentService {

}

#[tonic::async_trait]
impl Comments for CommentService{
        async fn create_comment(&self,request: Request<CommentRequest>,) -> Result<Response<CommentResponse>, Status> {
            println!("Got a request: {:?}", request);

            let req = request.into_inner();
            create_comment(req);
            let reply = CommentResponse {
                status: format!("Added new comment!").into(),
            };
            Ok(Response::new(reply))
        }
        async fn update_comment(&self,request: Request<CommentRequest>,) -> Result<Response<CommentResponse>, Status> {
            let reply = CommentResponse {
                status: format!("Updated comment!").into(),
            };
            let req = request.into_inner();
            update_comment(req);

            Ok(Response::new(reply))
        }
        async fn get_by_review(&self,request: Request<IdRequest>,) -> Result<Response<AllCommentsResponse>,Status> {
            let req = request.into_inner();
            let result = get_by_review(req.id);
            let comments =  result.iter().map(|comment_value | map_comment_to_comment_request(comment_value)).collect::<Vec<CommentRequest>>();
            let response = AllCommentsResponse {comments : comments};
            Ok(Response::new(response))
        }
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {

    // use self::schema::comments::dsl::*;

    let _connection = establish_connection();
    let addr = "[::1]:50055".parse()?;
    let comment_service = CommentService::default();

    Server::builder()
        .add_service(CommentsServer::new(comment_service))
        .serve(addr)
        .await?;

    Ok(())
}