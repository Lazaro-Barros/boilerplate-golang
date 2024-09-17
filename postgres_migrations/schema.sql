CREATE TABLE todo (
    uuid UUID PRIMARY KEY,               -- Campo 'uuid' do tipo UUID, chave primária
    title VARCHAR(255) NOT NULL,          -- Campo 'title', texto com até 255 caracteres, obrigatório
    description TEXT,                     -- Campo 'description', texto de tamanho variável
    done BOOLEAN NOT NULL DEFAULT false  -- Campo 'done', booleano com valor padrão como 'false', obrigatório
);
