const POST_AND_POINT = "post"

const postService = {
	getPostById: async (id, req) => {
		const responseDataPost = await fetch(`${req ? process.env.API_URL : process.env.API_URL_CLIENT}${POST_AND_POINT}`, {
			method: "POST",
			body: JSON.stringify({ id })
		})
		return await responseDataPost.json()
	}
}

export default postService
