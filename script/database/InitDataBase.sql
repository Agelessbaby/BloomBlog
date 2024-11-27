ALTER TABLE comment DROP FOREIGN KEY fk_comment_parent;
ALTER TABLE comment DROP FOREIGN KEY fk_comment_reply;
ALTER TABLE comment
    ADD CONSTRAINT fk_comment_parent
        FOREIGN KEY (parent_id) REFERENCES comment (id) ON DELETE CASCADE;
ALTER TABLE comment
    ADD CONSTRAINT fk_comment_reply
        FOREIGN KEY (reply_id) REFERENCES comment (id) ON DELETE CASCADE;
