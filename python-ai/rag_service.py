import os from supabase import create_client from openai import OpenAI

Initialize clients

supabase = create_client(os.getenv("SUPABASE_URL"), os.getenv("SUPABASE_KEY")) ai_client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))

def get_account_recommendation(receipt_text): # Generate embedding for receipt embedding = ai_client.embeddings.create( input=receipt_text, model="text-embedding-3-small" ).data[0].embedding

# Search Supabase for similar historical classifications
rpc_response = supabase.rpc(
    'match_expense_accounts',
    {'query_embedding': embedding, 'match_threshold': 0.8, 'match_count': 1}
).execute()

if rpc_response.data:
    return rpc_response.data[0]['account_code']
return "DEFAULT-EXPENSE-CODE"