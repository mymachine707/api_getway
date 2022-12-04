SELECT 
    ar.id,
    ar.title,
    ar.body,
    ar.created_at,
    ar.updated_at,
    ar.deleted_at,
    au.id,
    au.firstname,
    au.lastname,
    au.created_at,
    au.updated_at,
    au.deleted_at
 FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = 'cc9be5d9-aa7b-48a7-9bd5-737fd48f37f0';

INSERT INTO author (middlename)
VALUES ("sasd"+"adadad"+"adasdad");


update author set fullname=?;  where customer = ?


Select * from author WHERE ((firstname ILIKE '%' || 'a' || '%') OR (lastname ILIKE '%' || 'a' || '%') OR (middlename ILIKE '%' || 's' || '%') OR (fullname ILIKE '%' || 'a' || '%')) AND deleted_at is null ;


SELECT au.id, au.firstname, au.lastname, au.middlename, au.fullname, au.created_at, au.updated_at, au.deleted_at FROM author AS au;