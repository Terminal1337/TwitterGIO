tokens = open('tokens.txt').read().splitlines()


with open('auth_token.txt','a') as f:
    for i in tokens:
        token = i.split("	")[0]
        f.write(f"{token}\n") 