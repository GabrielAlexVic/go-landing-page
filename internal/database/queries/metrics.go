package queries

const GetTotalPageViews = `
    SELECT COUNT(*) FROM page_views
`

const GetUniqueVisitors = `
    SELECT COUNT(DISTINCT visitor_hash) FROM page_views
`

const GetTotalLeads = `
    SELECT COUNT(*) FROM leads
`

const GetTotalEvents = `
    SELECT COUNT(*) FROM events
`

const GetTopUtmSources = `
    SELECT COALESCE(utm_source, 'Direto / Nenhum') AS name, COUNT(*) AS count 
    FROM page_views 
    GROUP BY COALESCE(utm_source, 'Direto / Nenhum') 
    ORDER BY count DESC 
    LIMIT 5
`

const GetTopDevices = `
    SELECT COALESCE(device_type, 'Desconhecido') AS name, COUNT(*) AS count 
    FROM page_views 
    GROUP BY COALESCE(device_type, 'Desconhecido') 
    ORDER BY count DESC
`

const GetTopEvents = `
    SELECT COALESCE(event_label, event_type) AS name, COUNT(*) AS count 
    FROM events 
    GROUP BY COALESCE(event_label, event_type) 
    ORDER BY count DESC 
    LIMIT 5
`

const GetTopLeadSources = `
    SELECT COALESCE(utm_source, 'Direto / Nenhum') AS name, COUNT(*) AS count 
    FROM leads 
    GROUP BY COALESCE(utm_source, 'Direto / Nenhum') 
    ORDER BY count DESC 
    LIMIT 5
`

const GetDailyPageViews = `
    SELECT TO_CHAR(created_at, 'YYYY-MM-DD') AS name, COUNT(*) AS count 
    FROM page_views 
    WHERE created_at >= NOW() - INTERVAL '30 days' 
    GROUP BY TO_CHAR(created_at, 'YYYY-MM-DD') 
    ORDER BY name ASC
`

const GetDailyLeads = `
    SELECT TO_CHAR(created_at, 'YYYY-MM-DD') AS name, COUNT(*) AS count 
    FROM leads 
    WHERE created_at >= NOW() - INTERVAL '30 days' 
    GROUP BY TO_CHAR(created_at, 'YYYY-MM-DD') 
    ORDER BY name ASC
`

