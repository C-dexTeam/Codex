-- name: GetPLanguages :many
SELECT 
    pl.id, pl.language_id, pl.name, pl.description, pl.download_cmd, pl.compile_cmd, pl.image_path,
    pl.file_extention, pl.monaco_editor, pl.created_at
FROM 
    t_programming_languages as pl
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(language_id)::text IS NULL OR us.language_id = sqlc.narg(language_id)) AND
    (sqlc.narg(name)::text IS NULL OR name ILIKE '%' || sqlc.narg(name)::text || '%') AND
    (sqlc.narg(description)::text IS NULL OR description ILIKE '%' || sqlc.narg(description)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetPLanguage :one
SELECT 
    pl.id, pl.language_id, pl.name, pl.description, pl.download_cmd, pl.compile_cmd, pl.image_path,
    pl.file_extention, pl.monaco_editor, pl.created_at
FROM 
    t_programming_languages as pl
WHERE
    pl.id = @programming_language_id;

-- name: CreatePLanguage :exec
INSERT INTO
    t_programming_languages (language_id, name, description, download_cmd, compile_cmd, image_path, file_extention, monaco_editor)
VALUES
    (@language_id, @name, @description, @download_cmd, @compile_cmd, @image_path, @file_extention, @monaco_editor);

-- name: UpdatePLanguage :exec
UPDATE
    t_programming_languages
SET
    language_id =  COALESCE(sqlc.narg(language_id)::TEXT, language_id),
    name =  COALESCE(sqlc.narg(name)::TEXT, name),
    description =  COALESCE(sqlc.narg(description)::TEXT, description),
    download_cmd =  COALESCE(sqlc.narg(download_cmd)::TEXT, download_cmd),
    compile_cmd =  COALESCE(sqlc.narg(compile_cmd)::TEXT, compile_cmd),
    image_path =  COALESCE(sqlc.narg(image_path)::TEXT, image_path),
    file_extention =  COALESCE(sqlc.narg(file_extention)::TEXT, file_extention),
    monaco_editor =  COALESCE(sqlc.narg(monaco_editor)::TEXT, monaco_editor)
WHERE
    id = @programming_language_id;

-- name: DeletePLanguage :exec
DELETE FROM
    t_programming_languages
WHERE
    id = @programming_language_id;