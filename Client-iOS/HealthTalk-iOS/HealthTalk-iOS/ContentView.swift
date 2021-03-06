//
//  ContentView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/02/14.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        TabView(selection: /*@START_MENU_TOKEN@*//*@PLACEHOLDER=Selection@*/.constant(1)/*@END_MENU_TOKEN@*/) {
            HomeView().tabItem {
                VStack {
                    Image(systemName: "a")
                    Text("ホーム")
                }
            }.tag(1)
            Text("トーク画面").tabItem {
                VStack {
                    Image(systemName: "a")
                    Text("トーク")
                }
            }.tag(2)
            Text("健康記録画面").tabItem {
                VStack {
                    Image(systemName: "a")
                    Text("健康記録")
                }
            }.tag(3)
        }.edgesIgnoringSafeArea(.top)
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
