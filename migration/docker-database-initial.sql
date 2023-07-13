create table persons(
    id serial primary key,
    name varchar,
    description varchar
);

INSERT INTO persons(name, description) VALUES
('Paulo Barbosa da Silva', 'Paulo Barbosa da Silva (Sabará, 25 de janeiro de 1790, 28 de janeiro de 1868) foi mordomo-mor da Casa Imperial do Brasil e deputado pela então província de Minas Gerais. A sua participação na fundação da cidade de Petrópolis foi decisiva quando mobilizou o seu companheiro de arma, o engenheiro major Júlio Frederico Koeler.'),
('Júlio Frederico Koeler', 'Júlio Frederico Koeler, ou Julius Friedrich Koeler, (Mainz, Grão Ducado de Hessen - Darmstadt, Alemanha, 16 de junho de 1804 – Petrópolis, 1847) foi um militar e engenheiro teuto-brasileiro.');