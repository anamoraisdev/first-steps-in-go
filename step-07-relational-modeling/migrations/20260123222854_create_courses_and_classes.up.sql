CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()

);

CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    title TEXT NOT NULL,
    starts_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_course
        FOREIGN KEY (course_id)
        REFERENCES courses(id)
        ON DELETE CASCADE
);