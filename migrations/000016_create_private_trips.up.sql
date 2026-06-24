-- Trips run by private drivers under an active driver-parent match (StartTrip/EndTrip/UpdatePrivateTripStatus)
CREATE TABLE private_trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    driver_id UUID NOT NULL REFERENCES private_drivers(id),
    match_id UUID NOT NULL REFERENCES driver_parent_matches(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'in_progress', 'completed', 'cancelled')),
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_private_trips_driver_id ON private_trips(driver_id);
CREATE INDEX idx_private_trips_match_id ON private_trips(match_id);
CREATE INDEX idx_private_trips_status ON private_trips(status);
