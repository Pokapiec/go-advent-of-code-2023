
def is_game_possible(line: str) -> int:
    game_id, games = line.split(': ')
    game_id = int(game_id.split(" ")[1])
    max_colors = {
        'red': 12,
        'green': 13,
        'blue': 14,
    }

    for game in games.split('; '):
        color_counts = game.split(', ')
        for color_count in color_counts:
            count, color = color_count.split(' ')
            color = color.strip()
            if int(count) > max_colors[color]:
                return 0
    
    return game_id
            

with open('day2_data.txt') as f:
    _sum = 0
    for line in f:
        res = is_game_possible(line)
        if res:
            _sum += res
            print(res)

    print(_sum)