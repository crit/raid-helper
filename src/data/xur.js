export class Location {
    constructor(data = {}) {
        this.expires = data.expires || ""
        this.destinationName = data.destinationName || ""
        this.locationName = data.locationName || ""
    }

    expired() {
        if (this.expires === "") return true

        return new Date(this.expires) < new Date()
    }
}

export const xur = {
    async location() {
        return fetch('https://raid-help.critrussell.com/data/xur-location.json')
            .then(res => res.json())
            .then(data => new Location(data))
            .catch(err => console.error(err))
    }
}
