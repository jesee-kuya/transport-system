-- Tracks which students are assigned to a trip and their boarding status (OnboardStudent/ViewBoardedStudents)
CREATE TABLE trip_students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trip_id UUID NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES students(id),
    boarded BOOLEAN NOT NULL DEFAULT FALSE,
    boarded_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (trip_id, student_id)
);

CREATE INDEX idx_trip_students_trip_id ON trip_students(trip_id);
CREATE INDEX idx_trip_students_student_id ON trip_students(student_id);
