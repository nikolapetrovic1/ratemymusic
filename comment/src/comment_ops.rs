use crate::comment::CommentRequest;
use crate::models::{Comment, NewComment};
use crate::db::establish_connection;
use diesel::prelude::*;


pub fn create_comment(comment1: CommentRequest) {
    println!("Creating Comment: {:?}", comment);
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();

    let new_comment = NewComment {
        comment: &comment1.comment,
        user_id: comment1.user_id,
        review_id: comment1.review_id,
    };

    diesel::insert_into(comments)
        .values(new_comment)
        .execute(&connection)
        .expect("Error saving new comment");
}