-- Created when a school admin connects a private driver (ConnectPrivateDriver) or
-- when a private parent connects their driver to a school (ConnectWithSchool)
CREATE TABLE school_private_driver_connections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    school_id UUID NOT NULL REFERENCES schools(id),
    driver_id UUID NOT NULL REFERENCES private_drivers(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (school_id, driver_id)
);

CREATE INDEX idx_school_private_driver_connections_school_id ON school_private_driver_connections(school_id);
CREATE INDEX idx_school_private_driver_connections_driver_id ON school_private_driver_connections(driver_id);
