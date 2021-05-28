import { NgModule } from "@angular/core";

import { AppRoutingModule } from "./app-routing.module";
import { AppComponent } from "./app.component";
import { LoginComponent } from "./login/login.component";
import { NotFoundComponent } from "./not-found/not-found.component";
import { HomeComponent } from './home/home.component';

@NgModule({
	declarations: [AppComponent, LoginComponent, NotFoundComponent, HomeComponent],
	imports: [AppRoutingModule],
	bootstrap: [AppComponent],
})
export class AppModule {}
