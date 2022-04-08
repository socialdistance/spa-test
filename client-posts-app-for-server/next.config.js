/** @type {import('next').NextConfig} */
const nextConfig = {
	env: {
		API_URL: "http://server:8081/",
		API_URL_CLIENT: "http://localhost:8081/"
	},
	distDir: "build"
}

module.exports = nextConfig
