package org.hay

import org.junit.Test

class StandardOutputTest {

    @Test
    fun testStandardOutput() {
        var builder = ProcessBuilder()
        builder.command("/Users/tj/sourcecode/learn-rust/macros/target/release/macros")
        var process = builder.start()
        process.inputReader().use {
            for (i in it.readLines()) {
                println(i)
            }
        }
    }

}