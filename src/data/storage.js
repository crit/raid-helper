const namespace = "raid-helper"

const Storage = (table) => {
    table = `${namespace}:${table}`

    return {
        read() {
            try {
                return JSON.parse(localStorage.getItem(table)) || null
            } catch (e) {
                console.error(e)
                return null;
            }
        },

        write(data) {
            try {
                localStorage.setItem(table, JSON.stringify(data))
            } catch (e) {
                console.error(e)
            }
        },

        reset(data) {
            this.truncate()
            this.write(data)
        },

        truncate() {
            localStorage.removeItem(table)
        }
    }
}

export default Storage;
