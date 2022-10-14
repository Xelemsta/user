CREATE TABLE "user" (
  id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
  email varchar(250) NOT NULL UNIQUE,
  first_name varchar(250) NOT NULL,
  last_name varchar(250) NOT NULL,
  password varchar(250) NOT NULL,
  nickname varchar(250),
  country varchar(250),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  UNIQUE (first_name, last_name, email)
);
