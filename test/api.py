import requests

def call_api(url):
    try:
        response = requests.get(url)

        if response.status_code == 200:
            return response.json()
        else:
            print(f"Erreur : {response.status_code}")
            return None
    except Exception as e:
        print(f"Une erreur est survenue : {e}")
        return None

api_url = 'http://localhost:8080/api/contact'

while True:
    data = call_api(api_url)

    if data:
        print("Données reçues :", data)
    
    input("Appuyez sur Entrée pour appeler à nouveau l'API ou fermez le programme pour arrêter...")
