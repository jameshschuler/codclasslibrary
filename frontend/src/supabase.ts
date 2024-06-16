import { createClient } from '@supabase/supabase-js'

const supabaseProjectId = import.meta.env.VITE_SUPABASE_PROJECT_ID
const supabaseUrl = `https://${supabaseProjectId}.supabase.co`
const publicAnonKey = import.meta.env.VITE_SUPABASE_ANON_KEY

export const supabase = createClient(supabaseUrl, publicAnonKey)
