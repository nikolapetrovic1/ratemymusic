table! {
    comments (id) {
        id -> Int4,
        comment -> Text,
        user_id -> Nullable<Int4>,
        review_id -> Nullable<Int4>,
    }
}
