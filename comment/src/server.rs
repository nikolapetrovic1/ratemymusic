#[macro_use]
extern crate diesel;
extern crate dotenv;


// use crate::models::Comment;

// use self::diesel::prelude::*;
use db::establish_connection;
use tonic::{transport::Server, Request, Response, Status};
use comment::comments_server::{Comments,CommentsServer};
use comment::{CommentRequest,CommentResponse};
use crate::comment_ops::create_comment;
use crate::models::{NewComment};


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
                status: format!("Sent {} BTC to", 200).into(),
            };
            Ok(Response::new(reply))
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