import {getCurrentToken} from "../services/localStorage.service"
const POSTS_END_POINT = "posts"

const postsService = {
	getPostsByPage: async (idPostsPage, req) => {
		const responseDataPosts = await fetch(`${req ? process.env.API_URL : process.env.API_URL_CLIENT}${POSTS_END_POINT}/${idPostsPage}`)
		return await responseDataPosts.json()
	},
	getAllLength: async (req) => {
		const responseLength = await fetch(`${req ? process.env.API_URL : process.env.API_URL_CLIENT}${POSTS_END_POINT}`)
		const { count } = await responseLength.json()
		return count
	},
	getDataForSearch: async (body) => {
		const responseDataForSearch = await fetch(`${process.env.API_URL_CLIENT}${POSTS_END_POINT}/search`, {
			method: "POST",
			body: JSON.stringify(body)
		})
		return await responseDataForSearch.json()
	},
	updatePost: async (body) => {
		delete body.comments
		const responseDataUpdatePost = await fetch(`${process.env.API_URL_CLIENT}${POSTS_END_POINT}/update/${body.id}`, {
			method: "PUT",
			headers: {
				"Authorization": `Bearer ${getCurrentToken()}`,
				"Content-type": "application/json"
			},
			body: JSON.stringify(body)
		})
		return await responseDataUpdatePost.json()
	}
}

export default postsService
