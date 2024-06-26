import g4f
import sys

g4f.logging = False # enable logging
g4f.check_version = False # Disable automatic version checking


# Automatic selection of provider

"""# streamed completion
response = g4f.ChatCompletion.create(
    model="gpt-3.5-turbo",
    messages=[{"role": "user", "content": "Hello"}],
    stream=True,
)

for message in response:
    print(message, flush=True, end='')
"""
# normal response

def _get_answer(question):
    try:
        response = g4f.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": f"srictly respond yes | no: {question}"}],
        # proxy="http://ufbgpuipz2ll05x:q5pd9f49caq7eda@rp.proxyscrape.com:6060"
    )  # alterative model setting
        if "Yes: " in response:

            response = response.replace("Yes: ", "")
        return response
    except:
        return False
prompt = sys.argv[1]
print(_get_answer(f"strictly write me only a comment directly within 50 words to reply in a tweet related to {prompt}"))