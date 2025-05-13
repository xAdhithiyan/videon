CREATE TABLE IF NOT EXISTS "video" (
  id SERIAL PRIMARY KEY,
  userId INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  videoType VARCHAR(255) NOT NULL,
  TotalChunks INT NOT NULL,

  compressVideo     BOOLEAN,
  generateThumbnail BOOLEAN,
  transcodeVideo    BOOLEAN,
  addWaterMark      BOOLEAN,
  videoSummary      BOOLEAN,

  FOREIGN KEY (userId) REFERENCES "user"(id) ON DELETE CASCADE
);