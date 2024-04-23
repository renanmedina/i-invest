CREATE TABLE IF NOT EXISTS events(
  id SERIAL,
  event_name text NOT NULL,
  object_id varchar(255) NOT NULL,
  object_type text NOT NULL,
  event_data jsonb NULL DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL DEFAULT NULL
);

CREATE INDEX idx_event_event_name ON events(event_name);
CREATE INDEX idx_event_event_object_id ON events(object_id);
CREATE INDEX idx_event_event_object_type ON events(object_type);