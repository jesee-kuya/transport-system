-- Created when a private parent requests a match with a private driver (MatchWithDriver/MatchWithParent)
CREATE TABLE driver_parent_matches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    driver_id UUID NOT NULL REFERENCES private_drivers(id),
    parent_id UUID NOT NULL REFERENCES private_parents(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'active', 'rejected', 'ended')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (driver_id, parent_id)
);

CREATE INDEX idx_driver_parent_matches_driver_id ON driver_parent_matches(driver_id);
CREATE INDEX idx_driver_parent_matches_parent_id ON driver_parent_matches(parent_id);
