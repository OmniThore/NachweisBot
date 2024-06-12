CREATE TABLE Person
(
    PersonID int NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    Name VARCHAR(255) NOT NULL
);
-- Add a Buffer for AUTO_INCREMENT
ALTER TABLE Person AUTO_INCREMENT=100;
INSERT INTO Person (PersonID ,Name) VALUES (12,'Max Mustermann');


CREATE TABLE Nachweise (
    NachweisID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    PersonID int NOT NULL,
    Datum DATE NOT NULL DEFAULT CURRENT_DATE,
    Nachweis TEXT,
    FOREIGN KEY (PersonID) REFERENCES Person(PersonID) ON DELETE CASCADE ON UPDATE CASCADE
);
ALTER TABLE Nachweise AUTO_INCREMENT=100;

INSERT INTO Nachweise 
(NachweisID, PersonID, Nachweis) VALUES (1, 12,'Nachweis 1');
