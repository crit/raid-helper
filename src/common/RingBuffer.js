export default class RingBuffer {
    constructor(data = []) {
        this.resetAt = data.length - 1
        this.data = data
        this.cursor = 0
    }

    items() {
        return this.data
    }

    set(item) {
        this.data[this.cursor] = item
        this.next()
    }

    next() {
        let {cursor} = this

        cursor++
        if (cursor > this.resetAt) {
            cursor = 0
        }

        this.cursor = cursor
    }
}
