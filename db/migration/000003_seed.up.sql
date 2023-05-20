INSERT INTO category (name) VALUES ('Barbell'), ('Dumbbell'), ('Machine'), ('Bodyweight'), ('Cardio'), ('Cables');
INSERT INTO muscle_group (name) VALUES ('Arms'), ('Back'), ('Chest'), ('Legs'), ('Shoulders');

-- @todo add more
INSERT INTO stock_exercise (name, muscle_group, category) VALUES
('Barbell Bench Press', 'Chest', 'Barbell'),
('Incline Barbell Bench Press', 'Chest', 'Barbell'),
('Decline Barbell Bench Press', 'Chest', 'Barbell'),
('Dumbbell Bench Press', 'Chest', 'Dumbbell'),
('Incline Dumbbell Bench Press', 'Chest', 'Dumbbell'),
('Decline Dumbbell Bench Press', 'Chest', 'Dumbbell'),
('Machine Chest Fly', 'Chest', 'Machine'),
('Cable Chest Fly', 'Chest', 'Cables'),
('Incline Cable Chest Fly', 'Chest', 'Cables'),
('Decline Cable Chest Fly', 'Chest', 'Cables'),
('Chest Dips', 'Chest', 'Bodyweight'),
('Machine Chest Press', 'Chest', 'Machine'),
('Bent Over Barbell Row', 'Back', 'Barbell'),
('Deadlifts', 'Back', 'Barbell'),
('Pull-ups', 'Back', 'Bodyweight'),
('Lat Pulldown Cable', 'Back', 'Cables'),
('Inverted Row Machine', 'Back', 'Machine'),
('Bent Over Row Dumbbell', 'Back', 'Dumbbell'),
('Seated Row Cable', 'Back', 'Cables'),
('Reverse Fly Dumbbell', 'Back', 'Dumbbell'),
('Seated Overhead Press Dumbbell', 'Shoulders', 'Dumbbell'),
('Lateral Raise', 'Shoulders', 'Dumbbell'),
('Front Lateral Raise','Shoulders', 'Dumbbell'),
('Standing Overhead Barbell Press','Shoulders', 'Barbell'),
('Squat','Legs', 'Barbell'),
('Seated Leg Curl','Legs', 'Machine'),
('Seated Leg Curl Single Leg','Legs', 'Machine'),
('Bulgarian Split Squat','Legs', 'Dumbbell'),
('Leg Extension','Legs', 'Machine'),
('Leg Extension Single Leg','Legs', 'Machine');
