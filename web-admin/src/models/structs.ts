export interface Sports {
    id: number
    sport_name: string
}

export interface League {
    id: number
    league_name: string
    sport_id: number
}

export interface Teams {
    id: number
    league_id: number
    team_id: number
    team_name: string
}
