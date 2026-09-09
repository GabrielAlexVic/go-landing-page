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
