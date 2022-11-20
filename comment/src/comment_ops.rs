use crate::comment::CommentRequest;
use crate::models::{Comment, NewComment};
use crate::db::establish_connection;
use diesel::prelude::*;


pub fn create_comment(comment_req: CommentRequest) {
    println!("Creating Comment: {:?}", comment);
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();

    let new_comment = NewComment {
        comment: &comment_req.comment,
        user_id: comment_req.user_id,
        review_id: comment_req.review_id,
        reports : 0,
    };

    diesel::insert_into(comments)
        .values(&new_comment)
        .execute(&connection)
        .expect("Error saving new comment");

}

pub fn update_comment(comment_req: CommentRequest){
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();

    let comment_update = Comment {
        id: comment_req.id,
        comment: comment_req.comment,
        user_id: comment_req.user_id,
        review_id: comment_req.review_id,
        reports: comment_req.reports,
    };

    
    diesel::update(comments.find(comment_update.id))
        .set(&comment_update)
        .execute(&connection)
        .expect("Error updating");
}

pub fn get_by_review(id_value:i32) -> Vec<Comment>{
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();

    let results = comments.filter(review_id.eq(id_value))
    .load::<Comment>(&connection)
    .expect("Error loading comments");
    return results;
}

pub fn report_comment(id_value:i32){
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();
    diesel::update(comments.find(id_value))
    .set(reports.eq(reports + 1))
    .execute(&connection)
    .expect("Error reporting");
}


pub fn delete_comment(id_value:i32){
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();
    diesel::delete(comments.filter(id.eq(id_value))).execute(&connection).expect("Error deleting");
}

pub fn get_all_by_report_count() -> Vec<Comment>{
    use crate::schema::comments::dsl::*;

    let connection = establish_connection();

    let results = comments.order_by(reports.desc())
    .load::<Comment>(&connection)
    .expect("Error loading comments");
    return results;
}

pub fn get_by_user(connection: PgConnection, id_value:i32) -> Vec<Comment> {
    
    use crate::schema::comments::dsl::*;

    let results = comments.filter(user_id.eq(id_value))
    .load::<Comment>(&connection)
    .expect("Error loading comments");
    return results;
}