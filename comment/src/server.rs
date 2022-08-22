#[macro_use]
extern crate diesel;
extern crate dotenv;

use comment_ops::get_all_by_report_count;
use db::establish_connection;
use models::map_comment_to_comment_request;
use tonic::{transport::Server, Request, Response, Status};
use comment::comments_server::{Comments,CommentsServer};
use comment::{CommentRequest,CommentResponse,AllCommentsResponse,IdRequest,ReportRequest};
use crate::comment_ops::{create_comment,update_comment,get_by_review,report_comment,delete_comment};



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
        async fn report_comment(&self,request:Request<ReportRequest>,) -> Result<Response<CommentResponse>, Status> {
            let req = request.into_inner();
            report_comment(req.id);
            let response = CommentResponse { status: format!("Reported comment!").into(),};

            Ok(Response::new(response))
        }
        async fn delete_comment(&self,request:Request<IdRequest>,) -> Result<Response<CommentResponse>, Status> {
            let req = request.into_inner();
            delete_comment(req.id);
            let response = CommentResponse { status: format!("Deleted comment!").into(),};
            Ok(Response::new(response))
        }
        async fn get_all_by_report_count(&self,request:Request<IdRequest>,) -> Result<Response<AllCommentsResponse>,Status> {
            let result = get_all_by_report_count();
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