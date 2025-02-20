package queries

var AnimeList = `
	query($userId:Int,$type:MediaType,$perPage:Int){
		Page(perPage:$perPage){
			mediaList(userId:$userId,type:$type,status_in:[CURRENT,REPEATING],sort:UPDATED_TIME_DESC){
				media{
					id 
					type 
					status(version:2)format 
					episodes 
					bannerImage 
					title{
						userPreferred
					}
					coverImage{
						large
					}
					nextAiringEpisode{
						airingAt
						timeUntilAiring
						episode
					}
				}
			}
		}
	}`
