package com.example.golangmobile

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.golangmobile.databinding.ActivityMainBinding
import kotlinx.coroutines.launch

class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        lifecycleScope.launch {
            try {
                val strJson = api.Api.fetchJobs()

                binding.tvMain.text = strJson
                Log.d("Jobs", strJson)
            } catch (e: Exception) {
                binding.tvMain.text = e.localizedMessage
                Log.d("Jobs", e.localizedMessage)
            }

        }

//        println(hello.Hello.getJson)
//        binding.tvMain.text = hello.Hello.hello()
    }
}