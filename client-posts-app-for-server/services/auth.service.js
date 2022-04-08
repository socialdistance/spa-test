const AUTH_END_POINT = "login"

const authService = {
	signIn: async (payload) => {
		const responseDataAuth = await fetch(`${process.env.API_URL_CLIENT}${AUTH_END_POINT}`, {
			method: "POST",
			body: JSON.stringify(payload)
		})
		return await responseDataAuth.json()
	}
}

export default authService
