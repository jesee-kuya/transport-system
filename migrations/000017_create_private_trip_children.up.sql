-- Tracks children on a private trip with boarding and delivery confirmation
-- boarded: set by the driver (OnboardPrivateStudent/ConfirmStudentBoarding)
-- received: set by the parent when the child is handed back (ReceiveStudent)
CREATE TABLE private_trip_children (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    private_trip_id UUID NOT NULL REFERENCES private_trips(id) ON DELETE CASCADE,
    child_id UUID NOT NULL REFERENCES children(id),
    boarded BOOLEAN NOT NULL DEFAULT FALSE,
    boarded_at TIMESTAMPTZ,
    received BOOLEAN NOT NULL DEFAULT FALSE,
    received_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (private_trip_id, child_id)
);

CREATE INDEX idx_private_trip_children_trip_id ON private_trip_children(private_trip_id);
CREATE INDEX idx_private_trip_children_child_id ON private_trip_children(child_id);
