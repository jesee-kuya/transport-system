-- School bus trips managed by school drivers (StartTrip/EndTrip/UpdateTripStatus)
CREATE TABLE trips (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    school_id UUID NOT NULL REFERENCES schools(id),
    driver_id UUID NOT NULL REFERENCES school_drivers(id),
    bus_id UUID NOT NULL REFERENCES buses(id),
    trip_type VARCHAR(20) NOT NULL CHECK (trip_type IN ('pickup', 'dropoff')),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'in_progress', 'completed', 'cancelled')),
    started_at TIMESTAMPTZ,
    ended_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_trips_school_id ON trips(school_id);
CREATE INDEX idx_trips_driver_id ON trips(driver_id);
CREATE INDEX idx_trips_bus_id ON trips(bus_id);
CREATE INDEX idx_trips_status ON trips(status);
