create table public.attachments (
    name text not null,
    category text not null,
    id uuid not null default gen_random_uuid (),
    constraint attachments_pkey primary key (id)
) tablespace pg_default;
create table public.games (
    id uuid not null default gen_random_uuid (),
    name text not null,
    constraint Game_pkey primary key (id)
) tablespace pg_default;
create table public.load_attachment (
    loadout_id uuid not null,
    attachment_id uuid not null,
    constraint load_attachment_pkey primary key (loadout_id, attachment_id),
    constraint load_attachment_attachment_id_fkey foreign key (attachment_id) references attachments (id),
    constraint load_attachment_loadout_id_fkey foreign key (loadout_id) references loadouts (id)
) tablespace pg_default;
create table public.loadouts (
    id uuid not null default gen_random_uuid (),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone null,
    title character varying not null,
    source text null,
    source_url text null,
    weapon_name text not null,
    weapon_category text not null,
    created_by uuid not null,
    game_id uuid not null,
    attachments uuid [] null,
    constraint loadouts_pkey primary key (id),
    constraint loadouts_created_by_fkey foreign key (created_by) references auth.users (id),
    constraint loadouts_game_id_fkey foreign key (game_id) references games (id)
) tablespace pg_default;