data = split(strip(read(joinpath(@__DIR__, "./input.txt"), String)), '\n')

numbers = [filter(isdigit, collect(s)) for s in data]

tot = sum([parse(Int64,(n[1]*last(n))) for n in numbers])

println(tot)